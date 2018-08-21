import {Injectable} from '@angular/core';
import * as io from "socket.io-client";

const SOCKET_SERVER = '/api/secure/socket-io';

@Injectable()
export class SocketService {
    private socket: any;

    constructor() {
        this.connect();
    }

    connect() {
        this.socket = io({path: SOCKET_SERVER});
        this.addSocketListener('news', (res) => {
            console.log('news', res);
            this.send('event_sent', { test: "test" });
            this.send('chat message', "message1");
        });
    }

    send(event, data) {
        this.socket.emit(event, data);
    }

    disconnect() {
        this.socket.disconnect();
    }

    addSocketListener(event, callback) {
        this.socket.on(event, callback);
    }
}
