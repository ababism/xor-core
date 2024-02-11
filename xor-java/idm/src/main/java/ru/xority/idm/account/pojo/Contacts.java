package ru.xority.idm.account.pojo;

import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.Optional;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * @author foxleren
 */
public record Contacts(
        @JsonProperty("tg_login")
        Optional<String> tgLogin)
{
    public static Contacts getEmpty() {
        return new Contacts(
                Optional.empty()
        );
    }

    public String toJson(ObjectMapper objectMapper) {
        try {
            return objectMapper.writeValueAsString(this);
        } catch (JsonProcessingException e) {
            // FIXME add log
            throw new RuntimeException(e);
        }
    }

    public static Contacts fromResultSet(ResultSet rs) throws SQLException {
        return new Contacts(
                Optional.ofNullable(rs.getString("tg_login"))
        );
    }
}
