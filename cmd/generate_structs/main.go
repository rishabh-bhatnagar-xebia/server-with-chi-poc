package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const ModulesDirectory = "/home/rishabh-bhatnagar/xebia/poc/db_with_gorm/v2/modules/"

// mapping between module name and the list of tables the module depend upon
var ModuleDependencyMapping = map[string][]string{
	// "opaman":        {},
	// "project":       {},
	// "auth":          {},
	// "git":           {},
	// "store":         {},
	// "terraform":     {},
	// "rediscache":    {},
	// "audit":         {"resource_events"},
	// "k8s":           {"resources", "projects", "environments", "resource_types", "components", "modules"},
	// "modulefetcher": {"modules"},
	"secretmanager": {"bindings", "secret", "secret_type"},
}

func getDSN() string {
	return "host=localhost user=postgres password=password dbname=rb port=5432 sslmode=disable"
}

func getGenerator(dsn string) (*gen.Generator, error) {
	g := gen.NewGenerator(gen.Config{})
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}), &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: "secret_manager.", SingularTable: false}})
	if err != nil {
		return nil, err
	}
	g.UseDB(db)
	return g, nil
}

func getModulePath(moduleName string) string {
	return filepath.Join(ModulesDirectory, moduleName, "secretmanager")
}

func generateStructs(g *gen.Generator, moduleName string, moduleDependencies []string) {
	g.OutPath = "v2/out_temp"
	g.ModelPkgPath = "v2/model_temp"
	log.Println("writing module to", g.OutPath)
	for _, tableName := range moduleDependencies {
		structMeta := g.GenerateModel(tableName)
		g.ApplyBasic(structMeta)
	}
	g.Execute()
}

func GenerateStructs(dsn string, moduleDependencies map[string][]string) error {
	g, err := getGenerator(dsn)
	if err != nil {
		return err
	}

	for moduleName, dependencies := range moduleDependencies {
		generateStructs(g, moduleName, dependencies)
	}

	return nil
}

func main() {
	fmt.Println(os.Getenv("DSN"), "#################")
	if err := GenerateStructs(os.Getenv("DSN"), ModuleDependencyMapping); err != nil {
		panic(err)
	}
}
