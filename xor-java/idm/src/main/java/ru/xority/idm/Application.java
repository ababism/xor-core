package ru.xority.idm;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;

/**
 * @author foxleren
 */
@SpringBootApplication(exclude = {DataSourceAutoConfiguration.class}) // disable datasource auto ping
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
