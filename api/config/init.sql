CREATE TABLE creator (
    id serial NOT NULL UNIQUE,
    name VARCHAR(64),
    profession VARCHAR(64),
    focus VARCHAR(64),
    youtube VARCHAR(64),
    CONSTRAINT creator_pk PRIMARY KEY (id)
);

INSERT INTO creator(name, profession, focus, youtube) values
    ('Ben Awad', 'Software Consultant', 'Startup Projects, Programming Memes, React', 'https://www.youtube.com/c/BenAwad97'),
    ('Be A Better Dev', 'Senior Software Engineer', 'Amazon Web Services', 'https://www.youtube.com/c/BeABetterDev'),
    ('The Primeagen', 'CEO', 'VIM, Rust, Go', 'https://www.youtube.com/c/ThePrimeagen/'),
    ('Byte Byte Go', 'Software Architect', 'System Design', 'https://www.youtube.com/c/ByteByteGo');