package vault

import "github.com/keeper-security/keeper-sdk-golang/storage"

type RecordTypeScope int32

const (
	RecordTypeScope_Standard         RecordTypeScope = 0
	RecordTypeScope_User             RecordTypeScope = 1
	RecordTypeScope_Enterprise       RecordTypeScope = 2
	RecordTypeScope_Pam              RecordTypeScope = 3
	RecordTypeScope_PamConfiguration RecordTypeScope = 4
)

type IRecordTypeField interface {
	FieldType() string
	Label() string
	Required() bool
}

type IRecordType interface {
	Id() int64
	Name() string
	Scope() RecordTypeScope
	Description() string
	Fields() []IRecordTypeField
	storage.IUid[string]
}
