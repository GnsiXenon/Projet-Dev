CREATE TABLE User (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(50),
    mail VARCHAR(100),
    hash VARCHAR(255),
    score INT
);

CREATE TABLE Challenge (
    id INT PRIMARY KEY AUTO_INCREMENT,
    flag VARCHAR(150)
);

CREATE TABLE History (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_user INT,
    id_challenge INT,
    FOREIGN KEY (id_user) REFERENCES User(id) ON UPDATE CASCADE,
    FOREIGN KEY (id_challenge) REFERENCES Challenge(id) ON UPDATE CASCADE
);

INSERT INTO Challenge (flag) VALUES ("HaCoeur{w3_h0p3_y0u_l34rn3d_s0m3th1ng_4nd_th4t_w3b_c4n_b3_fun_853643}");
INSERT INTO Challenge (flag) VALUES ("HaCoeur{G00d_j0b_y0u_f0und_th3_fl4g}");
INSERT INTO Challenge (flag) VALUES ("HaCoeur{Georges-Guynemer_7}");
INSERT INTO Challenge (flag) VALUES ("HaCoeur{B0T_N0T_A1}");
INSERT INTO Challenge (flag) VALUES ("HaCoeur{R4ns0mw4r3_1s_n0t_c00l}");
