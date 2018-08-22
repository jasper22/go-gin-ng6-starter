import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {models} from "./shared/types/proto-types";
import {SocketService} from "./shared/socket.service";
import {AuthService} from "./shared/auth.service";
import User = models.User;

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    public users: User[];
    public currentUser = null;
    public username = 'admin';
    public password = '1234';
    public new_username = '';
    public new_password = '';

    constructor(private http: HttpClient,
                private socketService: SocketService,
                private authService: AuthService) {

        this.authService.currentUser.subscribe((res) => {
            this.currentUser = res;
        });

        this.socketService.addSocketListener('connect', (res) => {
            console.log('connect', res);
            // send ping
            this.socketService.send('ping', "ping_message1");
            // connect to room
            this.socketService.send("join_room", "room1");
            this.socketService.addSocketListener("connected", (res) => {
                console.log("Connected:", res);
            });
        });
    }

    public login() {
        this.authService.login(this.username, this.password).subscribe((result) => {
            if (result) {
                console.log('Logged in!', this.authService.isAuthenticated());
                this.getAllUsers();

            } else {
                console.log('Error!');
            }
        });
    }

    public logout() {
        this.authService.logout();
    }

    public getAllUsers() {
        this.http.get("api/admin/user")
            .subscribe((users: User[]) => {
                console.log(users);
                this.users = users;
            });
    }

    public createUser() {
        const user: User = new User({
            username: this.new_username,
            password: this.new_password,
        });
        console.log(user);
        this.http.post("api/admin/user/", user)
            .subscribe((usr: User) => {
                console.log(usr);
                this.getAllUsers();
            });
    }
}
