package main

func main() {

	// The Agent should Parse and read the following flags on startup. If the flag is not specified, the corresponding
	// Environment Variable is read, otherwise the default value is applied.

	/*

		{
			tls: false,
			certFile : "",
			keyFile: "",
			enrollmentToken: "",


		}


	*/

	// This main function will encompass the flow of the agent tasks.
	// 1. Enrollment (Saturday/Monaday deadline)

	// 1.1 The agent should check for the existence of a certificate (used to communicate with fleet manager)
	// 1.2 If a certificate exists, the agent is already enrolled and no further action is required.
	// 1.3 Otherwise the agent generate a new private key and certificate signing request (CSR) (if not generated
	// already)
	// 1.4 The agent will use the CSR to complete the enrollment process (call CompleteEnrollment)
	// 1.5 The fleet manager will sign the certificate and return the signed certificate
	// 1.6 The agent will save the certificate and uses it for further requests to agent service.

	// 2. DataSourceManager
	// 2.1 A singleton object that is created after the Enrollment workflow completes, that handles the functionality of the
	// agent.

	// Functions of the DataSourceManager Singleton

	/*
		- Start(ctx Context)
			starts the DataSourceManager with a context that is used to shutdown when
			the context is cancelled.

			Start(ctx Context) creates the following channels:
				- messages chan Message a channel that receives messages from data sources.
				- stats chan StatUpdate a channel that receives statistics update from DataWriter
				- configurationChanges chan AgentConfiguration a channnel that receives from FleetManagerClient

			Start(ctx Context) starts the following go routines:
				- Calls FleetManagerClient.Start in a new go routine
				- Calls DataWriter.Start in a go routine
				- A loop that waits for either:
					- AgentConfiguration from configurationChanges (call applyConfigurationChanges)
					- ctx.Done to gracfully shutdown

			Start(ctx Context) also create a WaitGroup that WaitForShutdown uses to wait for Graceful shutdown (waits for Fleetmanager, DataWriter and Datasource go routines to complete)

		- applyConfigurations(config: AgentConfiguration)
			applyConfigurations(config: AgentConfiguration)called when a new configuration is
			received. For each datasource configured in On or Forecast, create a data source, apply configuration, and
			call ProduceMessages. For other datasources already running, cancel their context to cause
			ProduceMessage to stop, then call cleanup.

		- WaitForShutdown()
			waits for the WaitGroup created by Start. The DataSourceManager is shutdown via
			a call to a ctx.Cancel (the context passed to the DataSourceManager.Start function).
	*/

}
