ALTER TABLE ta.users
ADD COLUMN full_name varchar(255);
ADD COLUMN status_verification boolean default false;
ADD COLUMN created_at timestamp;
ADD COLUMN updated_at timestamp;
ADD COLUMN deleted_at timestamp default null;
