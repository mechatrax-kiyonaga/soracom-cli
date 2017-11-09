package main

func generateTrunkCommands(templateDir, outputDir string) error {
	subCommandTemplate, err := openTemplateFile(templateDir, "trunk.gotmpl")
	if err != nil {
		return err
	}

	argsSlice := []commandArgs{
		{
			Use:                       "bills",
			Short:                     "cli.bills.summary",
			Long:                      "cli.bills.description",
			CommandVariableName:       "BillsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "bills.go",
		},
		{
			Use:                       "coupons",
			Short:                     "cli.coupons.summary",
			Long:                      "cli.coupons.description",
			CommandVariableName:       "CouponsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "coupons.go",
		},
		{
			Use:                       "credentials",
			Short:                     "cli.credentials.summary",
			Long:                      "cli.credentials.description",
			CommandVariableName:       "CredentialsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "credentials.go",
		},
		{
			Use:                       "data",
			Short:                     "cli.data.summary",
			Long:                      "cli.data.description",
			CommandVariableName:       "DataCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "data.go",
		},
		{
			Use:                       "devices",
			Short:                     "cli.devices.summary",
			Long:                      "cli.devices.description",
			CommandVariableName:       "DevicesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "devices.go",
		},
		{
			Use:                       "event-handlers",
			Short:                     "cli.event-handlers.summary",
			Long:                      "cli.event-handlers.description",
			CommandVariableName:       "EventHandlersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "event_handlers.go",
		},
		{
			Use:                       "files",
			Short:                     "cli.files.summary",
			Long:                      "cli.files.description",
			CommandVariableName:       "FilesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "files.go",
		},
		{
			Use:                       "groups",
			Short:                     "cli.groups.summary",
			Long:                      "cli.groups.description",
			CommandVariableName:       "GroupsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "groups.go",
		},
		{
			Use:                       "lora-devices",
			Short:                     "cli.lora-devices.summary",
			Long:                      "cli.lora-devices.description",
			CommandVariableName:       "LoraDevicesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_devices.go",
		},
		{
			Use:                       "lora-gateways",
			Short:                     "cli.lora-gateways.summary",
			Long:                      "cli.lora-gateways.description",
			CommandVariableName:       "LoraGatewaysCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_gateways.go",
		},
		{
			Use:                       "lora-network-sets",
			Short:                     "cli.lora-network-sets.summary",
			Long:                      "cli.lora-network-sets.description",
			CommandVariableName:       "LoraNetworkSetsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "lora_network_sets.go",
		},
		{
			Use:                       "operator",
			Short:                     "cli.operator.summary",
			Long:                      "cli.operator.description",
			CommandVariableName:       "OperatorCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "operator.go",
		},
		{
			Use:                       "auth-keys",
			Short:                     "cli.operator.auth-keys.summary",
			Long:                      "cli.operator.auth-keys.description",
			CommandVariableName:       "OperatorAuthKeysCmd",
			ParentCommandVariableName: "OperatorCmd",
			FileName:                  "operator_auth_keys.go",
		},
		{
			Use:                       "orders",
			Short:                     "cli.orders.summary",
			Long:                      "cli.orders.description",
			CommandVariableName:       "OrdersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "orders.go",
		},
		{
			Use:                       "payment-history",
			Short:                     "cli.payment-history.summary",
			Long:                      "cli.payment-history.description",
			CommandVariableName:       "PaymentHistoryCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_history.go",
		},
		{
			Use:                       "payment-methods",
			Short:                     "cli.payment-methods.summary",
			Long:                      "cli.payment-methods.description",
			CommandVariableName:       "PaymentMethodsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_methods.go",
		},
		{
			Use:                       "payment-statements",
			Short:                     "cli.payment-statements.summary",
			Long:                      "cli.payment-statements.description",
			CommandVariableName:       "PaymentStatementsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payment_statements.go",
		},
		{
			Use:                       "payer-information",
			Short:                     "cli.payer-information.summary",
			Long:                      "cli.payer-information.description",
			CommandVariableName:       "PayerInformationCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "payer_information.go",
		},
		{
			Use:                       "webpay",
			Short:                     "cli.payment-methods.webpay.summary",
			Long:                      "cli.payment-methods.webpay.description",
			CommandVariableName:       "PaymentMethodsWebpayCmd",
			ParentCommandVariableName: "PaymentMethodsCmd",
			FileName:                  "payment_methods_webpay.go",
		},
		{
			Use:                       "products",
			Short:                     "cli.products.summary",
			Long:                      "cli.products.description",
			CommandVariableName:       "ProductsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "products.go",
		},
		{
			Use:                       "roles",
			Short:                     "cli.roles.summary",
			Long:                      "cli.roles.description",
			CommandVariableName:       "RolesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "roles.go",
		},
		{
			Use:                       "shipping-addresses",
			Short:                     "cli.shipping-addresses.summary",
			Long:                      "cli.shipping-addresses.description",
			CommandVariableName:       "ShippingAddressesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "shipping_addresses.go",
		},
		{
			Use:                       "sigfox-devices",
			Short:                     "cli.sigfox-devices.summary",
			Long:                      "cli.sigfox-devices.description",
			CommandVariableName:       "SigfoxDevicesCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "sigfox_devices.go",
		},
		{
			Use:                       "stats",
			Short:                     "cli.stats.summary",
			Long:                      "cli.stats.description",
			CommandVariableName:       "StatsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "stats.go",
		},
		{
			Use:                       "air",
			Short:                     "cli.stats.air.summary",
			Long:                      "cli.stats.air.description",
			CommandVariableName:       "StatsAirCmd",
			ParentCommandVariableName: "StatsCmd",
			FileName:                  "stats_air.go",
		},
		{
			Use:                       "beam",
			Short:                     "cli.stats.beam.summary",
			Long:                      "cli.stats.beam.description",
			CommandVariableName:       "StatsBeamCmd",
			ParentCommandVariableName: "StatsCmd",
			FileName:                  "stats_beam.go",
		},
		{
			Use:                       "subscribers",
			Short:                     "cli.subscribers.summary",
			Long:                      "cli.subscribers.description",
			CommandVariableName:       "SubscribersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "subscribers.go",
		},
		{
			Use:                       "users",
			Short:                     "cli.users.summary",
			Long:                      "cli.users.description",
			CommandVariableName:       "UsersCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "users.go",
		},
		{
			Use:                       "auth-keys",
			Short:                     "cli.users.auth-keys.summary",
			Long:                      "cli.users.auth-keys.description",
			CommandVariableName:       "UsersAuthKeysCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_auth_keys.go",
		},
		{
			Use:                       "password",
			Short:                     "cli.users.password.summary",
			Long:                      "cli.users.password.description",
			CommandVariableName:       "UsersPasswordCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_password.go",
		},
		{
			Use:                       "permissions",
			Short:                     "cli.users.permissions.summary",
			Long:                      "cli.users.permissions.description",
			CommandVariableName:       "UsersPermissionsCmd",
			ParentCommandVariableName: "UsersCmd",
			FileName:                  "users_permissions.go",
		},
		{
			Use:                       "vpg",
			Short:                     "cli.vpg.summary",
			Long:                      "cli.vpg.description",
			CommandVariableName:       "VpgCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "vpg.go",
		},
		{
			Use:                       "logs",
			Short:                     "cli.logs.summary",
			Long:                      "cli.logs.description",
			CommandVariableName:       "LogsCmd",
			ParentCommandVariableName: "RootCmd",
			FileName:                  "logs.go",
		},
	}

	for _, args := range argsSlice {
		f, err := openOutputFile(outputDir, args.FileName)
		if err != nil {
			return err
		}
		err = subCommandTemplate.Execute(f, args)
		if err != nil {
			return err
		}
	}

	return nil
}
