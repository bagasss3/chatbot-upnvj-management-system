-- +goose Up

CREATE TABLE IF NOT EXISTS faculties (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100),
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

CREATE TABLE IF NOT EXISTS majors (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100),
  faculty_id VARCHAR(100) NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp NULL
);

ALTER TABLE majors ADD FOREIGN KEY (faculty_id) REFERENCES faculties(id) ON DELETE RESTRICT ON UPDATE CASCADE;

INSERT INTO faculties (id, name) VALUES ("1", 'Fakultas Ilmu Komputer');
INSERT INTO faculties (id, name) VALUES ("2", 'Fakultas Ekonomi dan Bisnis');
INSERT INTO faculties (id, name) VALUES ("3", 'Fakultas Ilmu Sosial dan Politik');
INSERT INTO faculties (id, name) VALUES ("4", 'Fakultas Hukum');
INSERT INTO faculties (id, name) VALUES ("5", 'Fakultas Kedokteran');
INSERT INTO faculties (id, name) VALUES ("6", 'Fakultas Teknik');
INSERT INTO faculties (id, name) VALUES ("7", 'Fakultas Ilmu Kesehatan');

INSERT INTO majors (id, name, faculty_id) VALUES ("1", 'Vokasi - Perbankan dan Keuangan', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("2", 'Vokasi - Akutansi', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("3", 'Sarjana - Manajemen', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("4", 'Sarjana - Akutansi', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("5", 'Sarjana - Ekonomi Pembangunan', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("6", 'Sarjana - Ekonomi Syariah', "2");
INSERT INTO majors (id, name, faculty_id) VALUES ("7", 'Magister - Manajemen', "2");

INSERT INTO majors (id, name, faculty_id) VALUES ("8", 'Sarjana - Ilmu Komunikasi', "3");
INSERT INTO majors (id, name, faculty_id) VALUES ("9", 'Sarjana - Hubungan Internasional', "3");
INSERT INTO majors (id, name, faculty_id) VALUES ("10", 'Sarjana - Ilmu Politik', "3");
INSERT INTO majors (id, name, faculty_id) VALUES ("11", 'Sarjana - Sains Informasi', "3");

INSERT INTO majors (id, name, faculty_id) VALUES ("12", 'Sarjana - Hukum', "4");
INSERT INTO majors (id, name, faculty_id) VALUES ("13", 'Magister - Hukum', "4");

INSERT INTO majors (id, name, faculty_id) VALUES ("14", 'Sarjana - Kedokteran', "5");
INSERT INTO majors (id, name, faculty_id) VALUES ("15", 'Sarjana - Farmasi', "5");
INSERT INTO majors (id, name, faculty_id) VALUES ("16", 'Profesi - Pendidikan Profesi Dokter', "5");

INSERT INTO majors (id, name, faculty_id) VALUES ("17", 'Sarjana - Teknik Mesin', "6");
INSERT INTO majors (id, name, faculty_id) VALUES ("18", 'Sarjana - Teknik Industri', "6");
INSERT INTO majors (id, name, faculty_id) VALUES ("19", 'Sarjana - Teknik Perkapalan', "6");
INSERT INTO majors (id, name, faculty_id) VALUES ("20", 'Sarjana - Teknik Elektro', "6");

INSERT INTO majors (id, name, faculty_id) VALUES ("21", 'Vokasi - Keperawatan', "7");
INSERT INTO majors (id, name, faculty_id) VALUES ("22", 'Vokasi - Fisioterapi', "7");
INSERT INTO majors (id, name, faculty_id) VALUES ("23", 'Sarjana - Keperawatan', "7");
INSERT INTO majors (id, name, faculty_id) VALUES ("24", 'Sarjana - Kesehatan Masyarakat', "7");
INSERT INTO majors (id, name, faculty_id) VALUES ("25", 'Sarjana - Gizi', "7");
INSERT INTO majors (id, name, faculty_id) VALUES ("26", 'Profesi - Pendidikan Profesi Ners', "7");

INSERT INTO majors (id, name, faculty_id) VALUES ("27", 'Vokasi - Sistem Informasi', "1");
INSERT INTO majors (id, name, faculty_id) VALUES ("28", 'Sarjana - Informatika', "1");
INSERT INTO majors (id, name, faculty_id) VALUES ("29", 'Sarjana - Sistem Informasi', "1");

-- +goose Down

DELETE FROM majors WHERE id IN ("1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26","27","28","29");
DELETE FROM faculties WHERE id IN ("1","2","3","4","5","6","7");
DROP TABLE IF EXISTS majors;
DROP TABLE IF EXISTS faculties;