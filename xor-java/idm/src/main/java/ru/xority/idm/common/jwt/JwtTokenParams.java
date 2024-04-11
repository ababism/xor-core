package ru.xority.idm.common.jwt;

import java.util.HashMap;
import java.util.Map;

import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class JwtTokenParams {
    private String subject;
    private Map<String, String> extraClaims = new HashMap<>();
}
