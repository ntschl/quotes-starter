DROP TABLE IF EXISTS quotes;

CREATE TABLE quotes (
    ID varchar(50) PRIMARY KEY,
    quote varchar(255) NOT NULL,
    author varchar(50) NOT NULL
);

INSERT INTO quotes (ID, quote, author)
VALUES 
    ('ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26', 'Clear is better than clever.', 'Ronald McDonald'),
    ('84ca5b5f-38f0-4e00-bcf5-ae916e887690', 'Empty string check!', 'Squidward Tentacles'),
    ('b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2', 'Don''t panic.', 'Oprah Winfrey'),
    ('99fae00d-c5d4-4575-ba59-7e79efaff603', 'A little copying is better than a little dependency.', 'Chris Pratt'),
    ('5441f417-1379-4997-80bc-e2eac7523133', 'The bigger the interface, the weaker the abstraction.', 'Mary Poppins'),
    ('fc27cfd6-8f29-437f-b951-0a527fa2f7d3', 'With the unsafe package there are no guarantees.', 'Rob Dyrdek'),
    ('f05da4ce-398c-48fb-9a54-009ec3304319', 'Reflection is never clear.', 'Bobby Hill'),
    ('f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4', 'Don''t just check errors, handle them gracefully.', 'Shrek'),
    ('1627de76-c799-4b18-80c7-6151baf0f585', 'Documentation is for users.', 'Hermione Granger'),
    ('7ee7ccc8-21f0-4bea-af55-97553cb0d4d4', 'Errors are values.', 'Clark Kent'),
    ('9d17a91b-3525-4bae-9a34-c4de8155767a', 'Make the zero value useful.', 'Drake'),
    ('dd815990-0875-48f4-bf78-98ff8397dbed', 'Channels orchestrate; mutexes serialize.', 'Yo-Yo Ma'),
    ('e3669a09-3d4b-4aec-8b51-a3b3412e0603', 'Don''t communicate by sharing memory, share memory by communicating.', 'Prince'),
    ('1a10287d-e83b-45c5-ba94-f290710da7eb', 'Concurrency is not parallelism.', 'Lao Tzu'),
    ('2c371688-f482-4c77-943e-89937da93d27', 'Design the architecture, name the components, document the details.', 'Tony the Tiger');
