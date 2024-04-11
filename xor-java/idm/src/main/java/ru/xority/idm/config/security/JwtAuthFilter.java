package ru.xority.idm.config.security;

import java.io.IOException;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.lang.NonNull;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.web.authentication.WebAuthenticationDetailsSource;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.filter.OncePerRequestFilter;

import ru.xority.idm.common.jwt.JwtService;
import ru.xority.idm.service.AccountService;

/**
 * @author foxleren
 */
@Component
@RequiredArgsConstructor
public class JwtAuthFilter extends OncePerRequestFilter {
    private static final Logger logger = LoggerFactory.getLogger(JwtAuthFilter.class);
    private static final String AUTH_HEADER = "Authorization";

    private final JwtService jwtService;
    private final AccountService accountService;

    @Override
    protected void doFilterInternal(
            @NonNull HttpServletRequest request,
            @NonNull HttpServletResponse response,
            @NonNull FilterChain filterChain) throws ServletException, IOException {
        final String authHeader = request.getHeader("Authorization");
        if (!StringUtils.hasText(authHeader) || !authHeader.startsWith("Bearer ")) {
            filterChain.doFilter(request, response);
            return;
        }
        final String jwt = authHeader.substring(7);
        final String email = jwtService.extractEmail(jwt);
        if (StringUtils.hasText(email)
                && SecurityContextHolder.getContext().getAuthentication() == null) {
            System.out.println("???");
            UserDetails userDetails = accountService.userDetailsService().loadUserByUsername(email);
            if (jwtService.isTokenValid(jwt, email)) {
                SecurityContext context = SecurityContextHolder.createEmptyContext();
                UsernamePasswordAuthenticationToken authToken = new UsernamePasswordAuthenticationToken(
                        userDetails,
                        null,
                        userDetails.getAuthorities()
                );
                authToken.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
                context.setAuthentication(authToken);
                SecurityContextHolder.setContext(context);
            }
        }
        filterChain.doFilter(request, response);
//        String authHeader = request.getHeader(AUTH_HEADER);
//        if (!StringUtils.hasText(authHeader)) {
//            filterChain.doFilter(request, response);
//            return;
//        }
//
//        if (!authHeader.startsWith("Bearer ")) {
//            logger.error("Auth header does not contain bearer token");
//            filterChain.doFilter(request, response);
//            return;
//        }
//
//        String jwtToken = authHeader.substring(7);
//        String email = jwtService.extractEmail(jwtToken);
//        if (!StringUtils.hasText(email)) {
//            logger.error("Email in jwt token is empty");
//            return;
//        }
//
//        if (SecurityContextHolder.getContext().getAuthentication() != null) {
//            filterChain.doFilter(request, response);
//            return;
//        }
//
//        Optional<AccountEntity> accountO = accountService.get(AccountFilter.activeByEmail(email));
//        if (accountO.isEmpty()) {
//            logger.error("Account with email={} is not found", email);
//            return;
//        }
//        AccountEntity account = accountO.get();
//
//        List<GrantedAuthority> authorities = new ArrayList<>();
//        authorities.add(new SimpleGrantedAuthority("ADMIN"));
//
//        System.out.println("!!!");
//        if (jwtService.isTokenValid(jwtToken, account.getEmail())) {
//            System.out.println("???");
//            SecurityContext context = SecurityContextHolder.createEmptyContext();
//            UsernamePasswordAuthenticationToken authToken = new UsernamePasswordAuthenticationToken(
//                    account.getEmail(),
//                    null,
//                    authorities
//            );
//            authToken.setDetails(new WebAuthenticationDetailsSource().buildDetails(request));
//            context.setAuthentication(authToken);
//            SecurityContextHolder.setContext(context);
//        }
//        filterChain.doFilter(request, response);
    }
}
