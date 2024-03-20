package ru.xority.idm.controller;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.controller.dto.AccessTokenResponse;
import ru.xority.idm.controller.dto.LogInRequest;
import ru.xority.idm.controller.dto.RegisterRequest;
import ru.xority.idm.service.AuthService;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/auth")
@RequiredArgsConstructor
public class AuthController {
    private final AuthService authService;

    @PostMapping("/register")
    public AccessTokenResponse register(@RequestBody @Valid RegisterRequest registerRequest) {
        String accessToken = authService.register(registerRequest.getEmail(), registerRequest.getPassword());
        return new AccessTokenResponse(accessToken);
    }

    @PostMapping("/log-in")
    public AccessTokenResponse logIn(@RequestBody @Valid LogInRequest logInRequest) {
        String accessToken = authService.logIn(logInRequest.getEmail(), logInRequest.getPassword());
        return new AccessTokenResponse(accessToken);
    }
}
