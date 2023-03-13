drop table fakultas;
drop table jurusan;

ALTER TABLE "jurusan"
    DROP FOREIGN KEY "jurusan_fakultas_fk";

drop table program_studi;

ALTER TABLE "program_studi"
    DROP FOREIGN KEY "program_studi_jurusan_fk";

drop table biodata_mahasiswa;

drop table email_mahasiswa;

ALTER TABLE "email_mahasiswa"
    DROP FOREIGN KEY "email_mahasiswa_mahasiswa_fk";

drop table data_akademik_mahasiswa;

ALTER TABLE "data_akademik_mahasiswa"
    DROP FOREIGN KEY "data_akademik_mahasiswa_mahasiswa_fk";
