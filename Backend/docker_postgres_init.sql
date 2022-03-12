CREATE EXTENSION cube;
CREATE EXTENSION earthdistance;

create table college_data(
    college_id      numeric     primary key,
    name            text 		not null,
    address         text  		not null,
    images          text[]   	not null,
    email           text  		not null,
    contact_no      text 		not null,
    website         text 		not null,
    vr_image        text[]   	not null,
    longitude       double precision not null,
    latitude        double precision not null,
    hostel          bool		not null,
    cutoff          real 	not null,
    fees            real		not null,
    institute_type  bool		not null,
    details         text 		not null,
    courses         text[]   	not null,
    state           text        not null,
    course_type     text        not null,
    deemed          bool        not null,
    document_tokens TSVECTOR    not null
);
