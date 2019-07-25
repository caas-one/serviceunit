package setup

// 读取CICD的应用，用于自动创建创建profile.yaml

// ServiceUnits defines the ServiceUnit Collections
var ServiceUnits = []string{"bank", "nocard", "account-pay", "yop", "custom", "configcenter", "accounting"}

// ServiceMap defines the mapping of Service and Unit
var ServiceMap = map[string][]string{
	"bank": []string{
		"bankrouter-component-hessian",
		"bankrouter-hessian",
		"bankchannel-hessian",
		"bankinterface-hessian",
		"bankinterface-nocard-hessian",
		"banktask-hessian",
		"bank-cpu-hessian",
		"bankunifiedquery-hessian",
		"bankcooper-hessian",
	},
	"nocard": []string{
		"nc-pay-hessian",
		"nc-config-hessian",
		"nc-boss",
		"nc-auth-hessian",
		"cwh-hessian",
		"foundation-hessian",
	},
	"account-pay": []string{
		"account-pay-hessian",
		"account-pay-mboss",
		"bac-hessian",
		"bac-app",
		"advance-boss",
		"mp-auth",
	},
	"yop": []string{
		"yop-oauth2-api",
		"yop-monitor-hessian",
		"yop-center",
		"yop-boss",
		"yop-hessian",
		"yop-notifier-hessian",
	},
	"custom": []string{
		"merchant-platform-boss",
	},
	"configcenter": []string{
		"configcenter-hessian",
		"configcenter-boss",
	},
	"accounting": []string{
		"accounting-hessian",
		"merchant-account-hessian",
		"account-hessian",
		"fee-hessian",
		"cs-process-hessian",
		"cs-model-hessian",
		"fee-boss",
		"cs-boss",
		"accounting-boss",
		"activity-account-hessian",
		"activity-account-boss",
		"accounting-query-hessian",
		"account-manage-hessian",
		"accountfront-manage-hessian",
		"accountingfront-boss",
		"accountingfront-hessian",
		"accountingfront-query-hessian",
		"merchant-accountfront-hessian",
	},
}
