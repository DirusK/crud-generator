package generators

// Generator main contract for different architectures.
type Generator interface {
	GenerateMigration() error
	GenerateIntegrationTests() error
	GenerateRepository() error
	GenerateDomain() error
	GenerateService() error
	GenerateHandler() error
}
