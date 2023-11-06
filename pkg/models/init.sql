    -- Users Table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id varchar(30) UNIQUE NOT NULL,
    user_email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Dropdown Table
CREATE TABLE dropdown (
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
    FOREIGN KEY (dropdown_id) REFERENCES dropdown(id)
);

-- Job Applications Table
CREATE TABLE job_applications (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    company VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    application_status VARCHAR(50),
    application_type VARCHAR(50),
    resume_id INT,
    cover_letter_id INT,
    posting_url VARCHAR(255),
    salary_range VARCHAR(100),
    contact_id INT,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(id)
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

INSERT INTO dropdown (text, table_type) VALUES ('Ready to apply', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Applied', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Followed-up', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Phone Screen', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Technical', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Findal Round/Onsite', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Offered', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Signed', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Rejected', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Declined', 'application_status');
INSERT INTO dropdown (text, table_type) VALUES ('Lost Contact', 'application_status');

INSERT INTO users (user_id, user_email) VALUES ('5', 'test@gmail.com');