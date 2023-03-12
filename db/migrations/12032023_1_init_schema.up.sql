create table fakultas
(
    id   varchar(16) not null primary key,
    nama varchar(50) not null
) engine = INNODB;

create table jurusan
(
    id          varchar(16)  not null,
    nama        varchar(100) not null,
    id_fakultas varchar(8)   not null,

    primary key (id, id_fakultas),

    constraint jurusan_fakultas_fk
        foreign key (id_fakultas) references fakultas (id)
            on update cascade on delete cascade

) engine = INNODB;

create table program_studi
(
    id         varchar(16)                   not null,
    nama       varchar(50)                   not null,
    jenjang    enum ('D3', 'S1', 'S2', 'S3') not null,
    id_jurusan varchar(16)                   not null,

    primary key (id, id_jurusan),

    constraint program_studi_jurusan_fk
        foreign key (id_jurusan) references jurusan (id)
            on update cascade on delete cascade

) engine = INNODB;

create table biodata_mahasiswa
(
    id            varchar(36)                                             not null primary key,
    nama          varchar(50)                                             not null,
    nik           varchar(16)                                             not null,
    agama         enum ('Islam', 'Kristen', 'Budha', 'Hindu', 'Konghucu') not null,
    jenis_kelamin enum ('L', 'P')                                         not null,
    tanggal_lahir date                                                    not null,
    alamat        text                                                    not null
) engine = INNODB;

create table email_mahasiswa
(
    id_mahasiswa varchar(36) not null,
    email        varchar(50) not null,

    primary key (id_mahasiswa, email),

    constraint email_mahasiswa_mahasiswa_fk
        foreign key (id_mahasiswa) references biodata_mahasiswa (id)
            on update cascade on delete cascade

) engine = INNODB;

create table data_akademik_mahasiswa
(
    id_mahasiswa varchar(36) not null primary key,
    id_prodi     varchar(36) not null,
    nim          varchar(16) not null,
    angkatan     int         not null,

    constraint data_akademik_mahasiswa_mahasiswa_fk
        foreign key (id_mahasiswa) references biodata_mahasiswa (id)
            on update cascade on delete cascade

) engine = INNODB;