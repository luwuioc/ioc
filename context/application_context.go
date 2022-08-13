package context

import (
	"github.com/luwuioc/ioc/beanfactory"
)

var _ beanfactory.BeanFactory = (ApplicationContext)(nil)

type ApplicationContext interface{}
