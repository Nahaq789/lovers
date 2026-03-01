package log

import (
	"lovers/internal/domain/models/expense/expenseid"
	"lovers/internal/domain/models/expense/expenselogid"
	"lovers/internal/domain/models/group/groupid"
	"lovers/internal/domain/models/user/userid"
	"lovers/internal/domain/models/valueobjects/createdat"
)

type ExpenseLog struct {
	expenseLogId expenselogid.ExpenseLogId
	expenseId    expenseid.ExpenseId
	groupId      groupid.GroupId
	userId       userid.UserId
	operation    string
	beforeData   string
	afterData    string
	createdAt    createdat.CreatedAt
}

func NewExpenseLog(
	expenseLogId expenselogid.ExpenseLogId,
	expenseId expenseid.ExpenseId,
	groupId groupid.GroupId,
	userId userid.UserId,
	operation string,
	beforeData string,
	afterData string,
	createdAt createdat.CreatedAt,
) *ExpenseLog {
	return &ExpenseLog{
		expenseLogId: expenseLogId,
		expenseId:    expenseId,
		groupId:      groupId,
		userId:       userId,
		operation:    operation,
		beforeData:   beforeData,
		afterData:    afterData,
		createdAt:    createdAt,
	}
}

func (el *ExpenseLog) GetExpenseLogId() expenselogid.ExpenseLogId {
	return el.expenseLogId
}

func (el *ExpenseLog) GetExpenseId() expenseid.ExpenseId {
	return el.expenseId
}

func (el *ExpenseLog) GetGroupId() groupid.GroupId {
	return el.groupId
}

func (el *ExpenseLog) GetUserId() userid.UserId {
	return el.userId
}

func (el *ExpenseLog) GetOperation() string {
	return el.operation
}

func (el *ExpenseLog) GetBeforeData() string {
	return el.beforeData
}

func (el *ExpenseLog) GetAfterData() string {
	return el.afterData
}

func (el *ExpenseLog) GetCreatedAt() createdat.CreatedAt {
	return el.createdAt
}
