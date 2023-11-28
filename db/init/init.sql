    -- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    google_id varchar(30) UNIQUE NOT NULL,
    user_email VARCHAR(255) UNIQUE NOT NULL,
    verified_email BOOLEAN,
    picture varchar(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    session_token UUID UNIQUE NOT NULL,
    user_id INT NOT NULL,
    expiration TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Dropdown Table
CREATE TABLE dropdowns (
    id SERIAL PRIMARY KEY,
    user_id INT,
    text VARCHAR(255) NOT NULL,
    table_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Tasks Table
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    task_status VARCHAR(50),
    task TEXT NOT NULL,
    dropdown_id INT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (dropdown_id) REFERENCES dropdowns(id)
);

-- Documents Table
CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    original_file_name VARCHAR(255) NOT NULL,
    path TEXT NOT NULL, -- system path
    notes TEXT,
    document_name VARCHAR(255) NOT NULL, -- name provided by user to help him/her find it
    type_of_document VARCHAR(255) NOT NULL, -- type such as resume, CV etc
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Job Applications Table
CREATE TABLE job_applications (
    id SERIAL PRIMARY KEY,
    date_applied TIMESTAMP WITH TIME ZONE,
    user_id INT NOT NULL,
    company VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    application_status_id INT,
    application_type_id INT,
    resume_id INT,
    cover_letter_id INT,
    posting_url VARCHAR(255),
    salary_low INT,
    salary_high INT,
    contact_id INT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (resume_id) REFERENCES documents(id),
    FOREIGN KEY (cover_letter_id) REFERENCES documents(id),
    FOREIGN KEY (application_status_id) REFERENCES dropdowns(id),
    FOREIGN KEY (application_type_id) REFERENCES dropdowns(id)
);



INSERT INTO dropdowns (text, table_type) VALUES ('Ready to apply', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Applied', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Followed-up', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Phone Screen', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Technical', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Final Round/Onsite', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Offered', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Signed', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Rejected', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Declined', 'application_status');
INSERT INTO dropdowns (text, table_type) VALUES ('Lost Contact', 'application_status');

INSERT INTO dropdowns (text, table_type) VALUES ('Traditional', 'application_type');
INSERT INTO dropdowns (text, table_type) VALUES ('Quick Apply', 'application_type');
INSERT INTO dropdowns (text, table_type) VALUES ('Referral', 'application_type');
INSERT INTO dropdowns (text, table_type) VALUES ('Inbound', 'application_type');
INSERT INTO dropdowns (text, table_type) VALUES ('Extra Effort', 'application_type');

INSERT INTO users (google_id, user_email, verified_email) VALUES ('5', 'test@gmail.com', 'false');
INSERT INTO sessions (user_id, session_token, expiration) VALUES ('1', '4d48ce83-3b36-4b82-b36b-327a8e2668f4', CURRENT_TIMESTAMP + INTERVAL '1 year');