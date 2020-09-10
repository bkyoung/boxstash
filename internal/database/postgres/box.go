package postgres

import (
	"boxstash/internal/boxstash/entities"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
	"time"
)

func (db *DB) CreateBox(ctx context.Context, box entities.Box) (int64, error) {
	var id int64

	var epoch time.Time
	if box.CreatedAt == epoch {
		box.CreatedAt = time.Now()
	}

	if box.UpdatedAt == epoch {
		box.UpdatedAt = time.Now()
	}

	tag := fmt.Sprintf("%s/%s", box.Username, box.Name)
	box.Tag = &tag

	err := db.QueryRow(ctx, `
        INSERT INTO
            boxes (
                name,
                user_id,
                is_private,
                created_at,
                updated_at,
                short_description,
                description,
                description_html,
                description_markdown,
                tag,
                downloads
            )
        VALUES(
            $1,
            (select id from users where username = $2),
            $3,
            $4,
            $5,
            $6,
            $7,
            $8,
            $9,
            $10,
            $11
        ) RETURNING id;
    `, box.Name,
		box.Username,
		box.Private,
		box.CreatedAt,
		box.UpdatedAt,
		box.ShortDescription,
		box.Description,
		box.DescriptionHTML,
		box.DescriptionMarkdown,
		box.Tag,
		box.Downloads).Scan(&id)
	if err != nil {
		db.Logger.Log("error", "error inserting new box record", "err", err, "component", "boxstash.internal.database.postgres.CreateBox")
		return id, err
	}
	db.Logger.Log("debug", "new box record inserted", "box_id", id, "component", "boxstash.internal.database.postgres.CreateBox")

	return id, nil
}

func (db *DB) ReadBoxByID(ctx context.Context, id int64) (entities.Box, error) {
	b := entities.Box{}
	err := db.QueryRow(ctx, `
        SELECT
            boxes.id,
            boxes.name,
            users.username,
            boxes.is_private,
            boxes.created_at,
            boxes.updated_at,
            boxes.short_description,
            boxes.description,
            boxes.description_html,
            boxes.description_markdown,
            boxes.tag,
            boxes.downloads
        FROM
            boxes
        JOIN
            users
        ON
            boxes.user_id = users.id
        WHERE
            boxes.id = $1;
        `, id).Scan(
		&b.ID,
		&b.Name,
		&b.Username,
		&b.Private,
		&b.CreatedAt,
		&b.UpdatedAt,
		&b.ShortDescription,
		&b.Description,
		&b.DescriptionHTML,
		&b.DescriptionMarkdown,
		&b.Tag,
		&b.Downloads,
	)
	if err != nil {
		db.Logger.Log("error", "error retreiving box record by id", "box_id", id, "err", err, "component",
			"boxstash.internal.database.postgres.ReadBoxByID")
		return entities.Box{}, err
	}
	db.Logger.Log("debug", "box record found", "box_id", b.ID, "box", b.Tag, "component", "boxstash.internal.database.postgres.ReadBoxByID")
	return b, nil
}

func (db *DB) ReadBoxByName(ctx context.Context, username, name string) (entities.Box, error) {
	box := entities.Box{}
	q := fmt.Sprintf(`
        SELECT
            boxes.id,
            boxes.name,
            users.username,
            boxes.is_private,
            boxes.created_at,
            boxes.updated_at,
            boxes.short_description,
            boxes.description_html,
            boxes.description_markdown,
            boxes.tag,
            boxes.downloads
        FROM
            boxes
        JOIN
            users
        ON
            boxes.user_id = users.id
        WHERE
            users.username = '%s' AND boxes.name = '%s';
    `, username, name)
	if err := pgxscan.Get(ctx, db, &box, q); err != nil {
		db.Logger.Log("error", "error retrieving box record", "box", fmt.Sprintf("%s/%s", username, name), "err", err, "component",
			"boxstash.internal.database.postgres.ReadBoxByName")
		return entities.Box{}, err
	}
	db.Logger.Log("debug", "box record found", "box_id", box.ID, "box", box.Tag, "component", "boxstash.internal.database.postgres.ReadBoxByName")
	return box, nil
}

func (db *DB) UpdateBox(ctx context.Context, updates map[string]interface{}) (int64, error) {
	id, err := db.getBoxID(ctx, updates["username"].(string), updates["name"].(string))
	if err != nil {
		tag := fmt.Sprintf("%s/%s", updates["username"], updates["name"])
		db.Logger.Log("error", "problem looking up existing box record", "box", tag, "err", err, "component",
			"boxstash.internal.database.postgres.UpdateBox")
		return id, err
	}

	c := []string{}
	delete(updates, "username")
	for k, v := range updates {
		switch v.(type) {
		case string:
			col := fmt.Sprintf("%s = '%s'", k, v)
			c = append(c, col)
		default:
			col := fmt.Sprintf("%s = %v", k, v)
			c = append(c, col)
		}
	}
	columns := strings.Join(c, ", ")
	condition := fmt.Sprintf("id = %d;", id)
	sqlQuery := fmt.Sprintf("UPDATE boxes SET %s WHERE %s", columns, condition)
	db.Logger.Log("debug", "sql", sqlQuery, "component", "boxstash.internal.database.postgres.UpdateBox")

	_, err = db.Exec(ctx, sqlQuery)
	if err != nil {
		db.Logger.Log("error", "error updating box record", "err", err, "component", "boxstash.internal.database.postgres.UpdateBox")
		return id, err
	}
	db.Logger.Log("debug", "box record updated", "box_id", id, "component", "boxstash.internal.database.postgres.UpdateBox")

	return id, nil
}

func (db *DB) DeleteBox(ctx context.Context, username, name string) (box entities.Box, err error) {
    box, err = db.ReadBoxByName(ctx, username, name);if err != nil {
        tag := fmt.Sprintf("%s/%s", username, name)
        db.Logger.Log("error", "problem deleting requested box record", "box", tag, "err", err, "component",
            "boxstash.internal.database.postgres.DeleteBox")
        return box, err
    }
    _, err = db.Exec(ctx, `DELETE from boxes WHERE id = $1`, box.ID);if err != nil {
        tag := fmt.Sprintf("%s/%s", username, name)
        db.Logger.Log("error", "problem deleting requested box record", "box", tag, "err", err, "component",
            "boxstash.internal.database.postgres.DeleteBox")
        return box, err
    }
	return box, nil
}

func (db *DB) getBoxID(ctx context.Context, username, name string) (id int64, err error) {
    err = db.QueryRow(ctx, `
        SELECT 
            boxes.id 
        FROM 
            boxes 
        JOIN 
            users 
        ON 
            boxes.user_id = users.id 
        WHERE 
            boxes.name = $1 and users.username = $2
    `, name, username).Scan(&id)
    if err != nil {
        tag := fmt.Sprintf("%s/%s", username, name)
        db.Logger.Log("error", "problem looking up existing box record", "box", tag, "err", err, "component",
            "boxstash.internal.database.postgres.UpdateBox")
    }
    return id, err
}