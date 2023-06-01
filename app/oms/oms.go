package oms

type omsAcl struct {
}

func (o omsAcl) GetSNodeNameList() (sNodeNameList []string) {
	//TODO implement me
	panic("implement me")
}

type OmsAcl interface {
	GetSNodeNameList() (sNodeNameList []string)
}

func NewOmsAcl() *omsAcl {
	return &omsAcl{}
}
