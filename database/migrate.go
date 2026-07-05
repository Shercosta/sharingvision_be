package database

import (
	"context"
	"log"

	"sharingvisionbe/ent"

	"entgo.io/ent/dialect/sql/schema"
)

func AutoMigrate(client *ent.Client) {
	err := client.Schema.Create(
		context.Background(),

		// Optional options
		schema.WithForeignKeys(true),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Migration completed")
}
