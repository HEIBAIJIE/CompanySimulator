package main

type Orgnization struct {
	members []Member
}

type OrgnizationRole int32

const (
	LEADER        OrgnizationRole = 0
	NORMAL_MEMBER OrgnizationRole = 1
)
