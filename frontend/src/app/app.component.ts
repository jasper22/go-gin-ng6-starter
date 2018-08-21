import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {models} from "./shared/types/proto-types";
import User = models.User;
import {SocketService} from "./shared/socket.service";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    public users: User[];

    constructor(private http: HttpClient,
                private socketService: SocketService) {
        this.getAllUsers();
        this.socketService.send("chat message", "message2");
    }

    public getAllUsers() {
        this.http.get("api/secure/user")
            .subscribe((users: User[]) => {
                this.users = users;
            });
    }

    public createUser() {
        let user: User = new User();
        user.Name = "User";
        user.Phone = "1234";
        this.http.post("api/secure/user", user)
            .subscribe((user: User) => {
                console.log(user);
                this.getAllUsers();
            });
    }
}
