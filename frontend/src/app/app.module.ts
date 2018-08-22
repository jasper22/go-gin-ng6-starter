import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppComponent} from './app.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from "@angular/common/http";
import {JsonInterceptor} from "./shared/json.interceptor";
import {SocketService} from "./shared/socket.service";
import {AuthService} from "./shared/auth.service";
import {RouterModule} from "@angular/router";
import {ROUTES} from "./app.routes";
import {JwtInterceptor} from "./shared/jwt.interceptor";
import {FormsModule} from "@angular/forms";

@NgModule({
    declarations: [
        AppComponent
    ],
    imports: [
        BrowserModule,
        HttpClientModule,
        FormsModule,
        RouterModule.forRoot(ROUTES)
    ],
    exports: [
        RouterModule
    ],
    providers: [
        AuthService,
        SocketService,
        {provide: HTTP_INTERCEPTORS, useClass: JsonInterceptor, multi: true},
        {provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true}
    ],
    bootstrap: [AppComponent]
})
export class AppModule {
}
