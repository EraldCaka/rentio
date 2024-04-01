

CREATE TABLE IF NOT EXISTS users (
                                     id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                     username VARCHAR UNIQUE,
                                     password VARCHAR,
                                     role INT
);

CREATE TABLE IF NOT EXISTS clients (
                                       id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                       username VARCHAR UNIQUE,
                                       password VARCHAR
);

CREATE TABLE IF NOT EXISTS buildings (
                                         id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                         client_id UUID,
                                         location VARCHAR,
                                         CONSTRAINT fk_clients FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE IF NOT EXISTS floors (
                                      id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                      building_id UUID,
                                      floor_number INT,
                                      CONSTRAINT fk_buildings FOREIGN KEY (building_id) REFERENCES buildings(id)
);
CREATE TABLE IF NOT EXISTS rooms (
                                     id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                     room_type INT,
                                     floor_id UUID,
                                     room_size INT,
                                     CONSTRAINT fk_floors FOREIGN KEY (floor_id) REFERENCES floors(id)
);


CREATE TABLE IF NOT EXISTS contracts (
                                         id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                         room_id UUID,
                                         user_id UUID,
                                         start_date TIMESTAMP,
                                         end_date TIMESTAMP,
                                         rent INT,
                                         status INT,
                                         CONSTRAINT fk_rooms FOREIGN KEY (room_id) REFERENCES rooms(id),
                                         CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS active_users (
                                            id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                            role INT,
                                            jwt_token VARCHAR
);
