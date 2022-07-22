ALTER TABLE necesidades."dependencia_necesidad"
ADD COLUMN IF NOT EXISTS dep_solicitante_id int NOT NULL;

ALTER TABLE necesidades."dependencia_necesidad"
ADD COLUMN IF NOT EXISTS dep_destino_id int NOT NULL;

ALTER TABLE necesidades."dependencia_necesidad"
ADD COLUMN IF NOT EXISTS dep_supervisor_id int;

ALTER TABLE necesidades."dependencia_necesidad"
ADD COLUMN IF NOT EXISTS rol_ordenador_gasto_id int;

COMMENT ON COLUMN necesidades."dependencia_necesidad".dep_solicitante_id IS E'Id de la dependencia solicitande de la necesidad (db: udistrital_core - esquema: oikos - tabla: dependencia).';
COMMENT ON COLUMN necesidades."dependencia_necesidad".dep_destino_id IS E'Id de la dependencia destino de la necesidad (db: udistrital_core - esquema: oikos - tabla: dependencia).';
COMMENT ON COLUMN necesidades."dependencia_necesidad".dep_supervisor_id IS E'Id de la dependencia del supervisor de la necesidad (db: udistrital_core - esquema: oikos - tabla: dependencia).';
COMMENT ON COLUMN necesidades."dependencia_necesidad".rol_ordenador_gasto_id IS E'Id de rol del ordenador del gasto (db: udistrital - esquema: core - tabla: ordenador_gasto).';
