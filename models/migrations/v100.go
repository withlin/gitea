package migrations

import (
	"code.gitea.io/gitea/models"
	"github.com/go-xorm/xorm"
)

func specifyWebhookSignatureType(x *xorm.Engine) error {
	return x.Sync2(new(models.Webhook))
}
