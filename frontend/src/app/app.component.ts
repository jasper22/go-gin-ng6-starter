import {Component} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {models} from "./shared/types/proto-types";
import User = models.User;
import {map} from "rxjs/operators";

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    public users: User[];

    constructor(private http: HttpClient) {
        this.getAllUsers()
    }

    public getAllUsers() {
        this.http.get("api/secure/user")
            .pipe(
                map((res: any) => JSON.parse(res))
            )
            .subscribe((users: User[]) => {
                this.users = users;
            });
    }

    public createUser() {
        let user: User = new User();
        user.Name = "Daniel";
        user.Phone = "1234";
        this.http.post("api/secure/user", user)
            .pipe(
                map((res: any) => JSON.parse(res))
            )
            .subscribe((user: User) => {
                console.log(user);
                this.getAllUsers();
            });
    }
}
