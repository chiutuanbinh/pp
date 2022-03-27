
CREATE TABLE IF NOT EXISTS legal_entities(
    id SERIAL PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS persons(
    id SERIAL PRIMARY KEY,
    name TEXT,
    age SMALLINT,
    legal_entity_id INT,
    image_url TEXT,
    CONSTRAINT fk_legal_entity
        FOREIGN KEY(legal_entity_id)
            REFERENCES legal_entities(id)
);
CREATE TABLE IF NOT EXISTS companies(
    id SERIAL PRIMARY KEY,
    legal_entity_id INT,
    name TEXT,
    CONSTRAINT fk_legal_entity
        FOREIGN KEY(legal_entity_id)
            REFERENCES legal_entities(id)
);

CREATE TABLE IF NOT EXISTS owns(
    id SERIAL PRIMARY KEY,
    legal_entity_id INT,
    company_id INT,
    CONSTRAINT fk_legal_entity
        FOREIGN KEY(legal_entity_id)
            REFERENCES legal_entities(ID),
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES companies(id)
);


CREATE TABLE IF NOT EXISTS projects(
    id SERIAL PRIMARY KEY,
    name TEXT,
    cost INT,
    expect_return INT
);
CREATE TABLE IF NOT EXISTS investments(
    id SERIAL PRIMARY KEY,
    project_id INT,
    company_id INT,
    investment_value INT,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES companies(id),
    CONSTRAINT fk_project
        FOREIGN KEY(project_id)
            REFERENCES projects(id)
);


CREATE TABLE IF NOT EXISTS industry_categories(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS industries(
    id SERIAL PRIMARY KEY,
    name TEXT,
    industry_category_id INT,
    CONSTRAINT fk_industry_category
        FOREIGN KEY(industry_category_id)
            REFERENCES industry_categories(id)
);

CREATE TABLE IF NOT EXISTS operating_in(
    id SERIAL PRIMARY KEY,
    company_id INT,
    industry_id INT ,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES companies(id),
    CONSTRAINT fk_industry
        FOREIGN KEY(industry_id)
            REFERENCES industries(id)
            
);


CREATE TABLE IF NOT EXISTS industry_relation_types(
    id SERIAL PRIMARY KEY,
    name TEXT,
    coefficient NUMERIC
);



CREATE TABLE IF NOT EXISTS indsutry_relations(
    id SERIAL PRIMARY KEY,
    from_company_id INT,
    to_company_id INT,
    industry_relation_type_id INT,
    CONSTRAINT fk_from_company
        FOREIGN KEY(from_company_id)
            REFERENCES companies(id),
    CONSTRAINT fk_to_company
        FOREIGN KEY(to_company_id)
            REFERENCES companies(id),
    CONSTRAINT fk_industry_relation_type
        FOREIGN KEY(industry_relation_type_id)
            REFERENCES industry_relation_types(id)
);


CREATE TABLE IF NOT EXISTS financial_statement_types(
    id SERIAL PRIMARY KEY,
    name TEXT
);


CREATE TABLE IF NOT EXISTS financial_statements(
    id SERIAL PRIMARY KEY,
    company_id INT,
    name TEXT,
    fiscal_year INT,
    fiscal_period TEXT,
    start_date BIGINT ,
    end_date BIGINT ,
    financial_statement_type_id INT,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES companies(id),
    CONSTRAINT fk_financial_statement_type
        FOREIGN KEY(financial_statement_type_id)
            REFERENCES financial_statement_types(id)
);


CREATE TABLE IF NOT EXISTS financial_statement_line_types(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS financial_statement_lines(
    id SERIAL PRIMARY KEY,
    name TEXT,
    Description TEXT,
    Value INT,
    financial_statement_line_type_id INT,
    CONSTRAINT fk_financial_statement_line_type
        FOREIGN KEY(financial_statement_line_type_id)
            REFERENCES financial_statement_line_types(ID)
);


CREATE TABLE IF NOT EXISTS financial_statement_line_sequences(
    id SERIAL PRIMARY KEY,
    financial_statement_line_id INT,
    financial_statement_id INT,
    sequence INT,
    CONSTRAINT fk_financial_statement_line
        FOREIGN KEY(financial_statement_line_id)
            REFERENCES financial_statement_lines(ID),
    CONSTRAINT fk_financial_statement   
        FOREIGN KEY(financial_statement_id)
            REFERENCES financial_statements(ID)
);


CREATE TABLE IF NOT EXISTS exchanges(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE IF NOT EXISTS stocks (
    id SERIAL PRIMARY KEY,
    code TEXT,
    exchange_id INT,
    CONSTRAINT fk_exchange
        FOREIGN KEY(exchange_id)
            REFERENCES exchanges(ID)
);

CREATE TABLE IF NOT EXISTS issues(
    id SERIAL PRIMARY KEY,
    stock_id INT,
    company_id INT,
    date BIGINT ,
    amount INT,
    initial_price INT,
    CONSTRAINT fk_company
        FOREIGN KEY(company_id)
            REFERENCES companies(ID),
    CONSTRAINT fk_stock
        FOREIGN KEY(stock_id)
            REFERENCES stocks(id)
);


CREATE TABLE IF NOT EXISTS stock_prices (
    id SERIAL PRIMARY KEY,
    date BIGINT ,
    opening_price INT,
    closing_price INT,
    highest INT,
    lowest INT,
    stock_id INT,
    CONSTRAINT fk_stock
        FOREIGN KEY(stock_id)
            REFERENCES stocks(id)
);

