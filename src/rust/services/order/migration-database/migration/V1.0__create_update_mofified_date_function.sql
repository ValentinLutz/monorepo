CREATE FUNCTION update_modified_date() RETURNS TRIGGER
    LANGUAGE plpgsql AS
$$
BEGIN
    NEW.modified_date := now();
    RETURN NEW;
END;
$$;