CREATE TABLE editions (
                 id serial not null,
                 
                 title varchar(15) not null,
                 cover varchar(300) not null,

                 articles integer not null,
                 ecohustles integer not null,

                 PRIMARY KEY  (title)
            );


CREATE TABLE articles (
				 id serial not null,
                 count integer not null,
                 
                 edition varchar(15) not null,
                 
                 title varchar(100) not null,
                 author varchar(200) not null,
                 email varchar(200) not null,
                 post varchar(65535) not null,

                 gravatar varchar(200) not null,

                 PRIMARY KEY  (id)
            );

CREATE TABLE editorial (
                edition varchar(15) unique,

                title varchar(100) not null,
                author varchar(200) not null,
                email varchar(200) not null,
                post varchar(65535) not null,

                gravatar varchar(200) not null,

                PRIMARY KEY (edition)
            );

CREATE TABLE featured (
				edition varchar(15) unique,

				title varchar(100) not null,
				author varchar(200) not null,
				email varchar(200) not null,
				post varchar(65535) not null,

				gravatar varchar(200) not null,

				PRIMARY KEY (edition)
			);


CREATE TABLE ecohustles (
                 id serial not null,
                 count integer not null,
                 
                 edition varchar(15) not null,

                 image varchar(300) not null,
                 caption varchar(400) not null,

                 PRIMARY KEY  (id)
            );