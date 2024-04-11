package ru.xority.idm.api.http;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.annotation.AuthenticationPrincipal;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.api.http.dto.AccessTokenResponse;
import ru.xority.idm.api.http.dto.LogInRequest;
import ru.xority.idm.api.http.dto.RegisterRequest;
import ru.xority.idm.mapper.AccountMapper;
import ru.xority.idm.service.AccountService;
import ru.xority.idm.service.AuthService;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/auth")
@RequiredArgsConstructor
public class AuthController {
    private final AuthService authService;
    private final AccountService accountService;
    private final AccountMapper accountMapper;

    @PostMapping("/register")
    public AccessTokenResponse register(@RequestBody @Valid RegisterRequest registerRequest) {
        String accessToken = authService.register(registerRequest.getEmail(), registerRequest.getPassword());
        return new AccessTokenResponse(accessToken);
    }

    @PostMapping("/log-in")
    public AccessTokenResponse logIn(@AuthenticationPrincipal UserDetails user,
                                     @RequestBody @Valid LogInRequest logInRequest) {
//        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
//        System.out.println(user.getUsername());
        String accessToken = authService.logIn(logInRequest.getEmail(), logInRequest.getPassword());
        return new AccessTokenResponse(accessToken);
    }
}
