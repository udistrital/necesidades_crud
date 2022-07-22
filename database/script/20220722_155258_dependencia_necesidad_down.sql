ALTER TABLE necesidades."dependencia_necesidad"
DROP COLUMN IF EXISTS dep_solicitante_id;

ALTER TABLE necesidades."dependencia_necesidad"
DROP COLUMN IF EXISTS dep_destino_id;

ALTER TABLE necesidades."dependencia_necesidad"
DROP COLUMN IF EXISTS dep_supervisor_id;

ALTER TABLE necesidades."dependencia_necesidad"
DROP COLUMN IF EXISTS rol_ordenador_gasto_id;
