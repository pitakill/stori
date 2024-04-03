package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"stori/internal/adapter/config"
	"stori/internal/core/domain"
	"stori/internal/core/port"

	"github.com/aws/aws-sdk-go-v2/aws"
	configAWS "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type data struct {
	User               string  `json:"user"`
	Total              float64 `json:"total"`
	TotalAverageDebit  float64 `json:"totalAverageDebit"`
	TotalAverageCredit float64 `json:"totalAverageCredit"`
	Months             string  `json:"months"`
}

const montsTemplate = `
{{range $key, $value := .TransactionByMonth}}
  <tr>
    <td>
      {{$key}}
    </td>
    <td class="alignright">
      {{$value}}
    </td>
  </tr>
{{end}}
`

type EmailService struct {
	repository port.UserRepository
	config     *config.Container
}

func NewSesEmail(repository port.UserRepository, config *config.Container) *EmailService {
	return &EmailService{
		repository,
		config,
	}
}

func (es *EmailService) SendEmail(ctx context.Context, email *domain.Email) error {
	user, err := es.repository.GetUserByAccountID(ctx, email.AccountID)
	if err != nil {
		return err
	}

	config, err := configAWS.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	client := ses.NewFromConfig(config)

	tpl, err := template.New("months").Parse(montsTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, email.Summary)
	if err != nil {
		return err
	}

	templateData := &data{
		User:               fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		Total:              email.Summary.Total.InexactFloat64(),
		TotalAverageDebit:  email.Summary.TotalAverageDebit.InexactFloat64(),
		TotalAverageCredit: email.Summary.TotalAverageCredit.InexactFloat64(),
		Months:             buf.String(),
	}

	b, err := json.Marshal(templateData)
	if err != nil {
		return err
	}

	_, err = client.SendTemplatedEmail(ctx, &ses.SendTemplatedEmailInput{
		Source:       aws.String(es.config.AWS.SES.EmailFrom),
		Template:     aws.String(es.config.AWS.SES.TemplateName),
		TemplateData: aws.String(string(b)),
		Destination: &types.Destination{
			ToAddresses: []string{user.Email},
		},
	})

	return err
}
