{{- /*gotype: crud-generator-gui/internal/generators/vipcoin/models.Entity*/ -}}
-- +migrate Up

{{.MigrationCreateExtensions}}
{{.MigrationCreateTypes}}

create table if not exists {{.TableName}}
(
    {{.MigrationTableFields}}
    created_at timestamp    default now()                     not null,
    updated_at timestamp    default now()                     not null
);

-- +migrate StatementBegin
create or replace function trigger_set_updated_at() returns trigger as
$$
begin
    NEW.updated_at = now();
    return NEW;
end
$$ language plpgsql;

create trigger {{.TableName}}_set_updated_at
    before update
    on {{.TableName}}
    for each row
execute procedure trigger_set_updated_at();
-- +migrate StatementEnd

-- +migrate Down

drop function if exists {{.TableName}}_set_updated_at cascade;
drop table {{.TableName}};
{{.MigrationDropTypes}}
{{.MigrationDropExtensions}}
