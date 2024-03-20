package ru.xority.idm.common.jwt;

/**
 * @author foxleren
 */
public interface JwtService {
    String extractEmail(String token);

    String generateToken(JwtTokenParams jwtTokenParams);

    boolean isTokenValid(String token, String email);
}
