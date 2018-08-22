import {Injectable} from '@angular/core';
import { Observable, ReplaySubject } from 'rxjs';
import {map} from 'rxjs/operators';
import {HttpClient} from '@angular/common/http';
import {Router} from '@angular/router';
import {models} from './types/proto-types';
import User = models.User;

@Injectable()
export class AuthService {
    public token: string;
    private baseUrl = 'api/public/auth';
    private currentUser$: ReplaySubject<User> = new ReplaySubject<User>(1);

    constructor(private http: HttpClient,
                private router: Router) {
        if (localStorage.getItem('jwt')) {
            this.token = localStorage.getItem('jwt');
            this.currentUser$.next(this.getUserInfoFromToken());
        }
    }

    public get currentUser(): ReplaySubject<User> {
        return this.currentUser$;
    }

    public set currentUser(value: ReplaySubject<User>) {
        this.currentUser$ = value;
    }

    public login(username: string, password: string): Observable<boolean> {
        return this.http.post(`${this.baseUrl}/login`, {username: username, password: password})
            .pipe(
                map((response: any) => {
                    // login successful if there's a jwt token in the response
                    const token = response && response.token;
                    if (token) {
                        this.token = token;
                        localStorage.setItem('jwt', token);
                        this.currentUser$.next(this.getUserInfoFromToken());
                        return true;
                    } else {
                        return false;
                    }
                })
            );
    }

    public signup(username: string, password: string): Observable<boolean> {
        return this.http.post(`${this.baseUrl}/signup`, {username: username, password: password})
            .pipe(
                map((response: any) => {
                    // login successful if there's a jwt token in the response
                    const token = response && response.token;
                    if (token) {
                        this.token = token;
                        localStorage.setItem('email-jwt', token);
                        this.currentUser$.next(this.getUserInfoFromToken());
                        return true;
                    } else {
                        return false;
                    }
                })
            );
    }

    public logout(): void {
        this.token = null;
        localStorage.removeItem('jwt');
        this.currentUser$.next(null);
        this.router.navigateByUrl('/');
    }

    public isAuthenticated() {
        const token = localStorage.getItem('jwt');
        if (token) {
            if (new Date() > new Date(this.parseToken(token).exp * 1000)) {
                localStorage.removeItem('jwt');
                this.currentUser$.next(null);
                return false;
            }
            return true;
        } else {
            return false;
        }
    }

    public getUserInfoFromToken(): User {
        const token = localStorage.getItem('jwt');
        if (token) {
            return this.parseToken(token).user;
        }
    }

    private parseToken(token) {
        const base64Url = token.split('.')[1];
        const base64 = base64Url.replace('-', '+').replace('_', '/');
        return JSON.parse(window.atob(base64));
    }
}
