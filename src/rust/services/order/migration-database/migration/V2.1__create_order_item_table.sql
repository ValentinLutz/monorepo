CREATE TABLE order_service.order_item
(
    order_item_id INT GENERATED ALWAYS AS IDENTITY,
    creation_date TIMESTAMPTZ NOT NULL,
    modified_date TIMESTAMPTZ NOT NULL DEFAULT now(),
    order_id      VARCHAR     NOT NULL,
    item_name     VARCHAR     NOT NULL,
    PRIMARY KEY (order_item_id),
    CONSTRAINT order_id FOREIGN KEY (order_id) REFERENCES order_service.order (order_id)
);

CREATE TRIGGER update_order_item_modified_date
    BEFORE UPDATE
    ON order_service.order
    FOR EACH ROW
EXECUTE PROCEDURE order_service.update_modified_date();
