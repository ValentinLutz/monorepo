CREATE TABLE order_service.order
(
    order_id      VARCHAR     NOT NULL UNIQUE,
    creation_date TIMESTAMPTZ NOT NULL,
    modified_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    workflow      VARCHAR     NOT NULL,
    order_status  VARCHAR     NOT NULL,
    PRIMARY KEY (order_id)
);

CREATE TRIGGER update_order_modified_date
    BEFORE UPDATE
    ON order_service.order
    FOR EACH ROW
EXECUTE PROCEDURE order_service.update_modified_date();
