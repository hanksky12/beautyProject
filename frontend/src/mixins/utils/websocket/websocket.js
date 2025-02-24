import {CreateRequestMixin} from "@/mixins/utils/handle-request/create-request/createRequest";
import {PrintToastMixin} from "@/mixins/utils/send-request-to-post-processing/printToast";
import {ResponseData} from "@/store/modules/ajax";
export const WebSocketMixin = {
    data() {
        return {
            webSockets: {},
        };
    },
    mixins: [
        CreateRequestMixin,
        PrintToastMixin,
    ],
    methods: {
        startWsConnect(wsName, requestData, alreadyExistsStr, onOpenFunc, onMsgFunc, onErrorFunc, onCloseFunc ) {
            console.log(`Connecting to WebSocket for ${wsName}...`);
            if (this.webSockets[wsName]) {
                console.warn(`WebSocket for ${wsName} already exists.`);
                this.printToast(new ResponseData({}, alreadyExistsStr, 201,true));
                return;
            }

            const request = this.createRequest(requestData, "ws");
            const webSocket= new WebSocket(request.url);
            this.defEvent(webSocket, onOpenFunc, onMsgFunc, onErrorFunc, wsName, onCloseFunc);
            this.webSockets[wsName] = webSocket;
        },
        defEvent(webSocket, onOpenFunc, onMsgFunc, onErrorFunc, wsName, onCloseFunc) {
            webSocket.onopen = () => {
                console.log("WebSocket connection established.");
                const startStr = onOpenFunc();
                this.printToast(new ResponseData(startStr, {}, 200, true));
            };

            webSocket.onmessage = (event) => {
                console.log("Message received from WebSocket:", event.data);
                const msgStr = onMsgFunc(event.data);
                this.printToast(new ResponseData(msgStr, {}, 200, true));
            };

            webSocket.onerror = (error) => {
                console.error("WebSocket error:", error);
                const errorStr = onErrorFunc(error);
                this.printToast(new ResponseData({}, errorStr, 0, false));
            };

            webSocket.onclose = (event) => {
                console.log("WebSocket connection closed:", event);
                delete this.webSockets[wsName]; // 清理對應的 WebSocket
                const stopStr = onCloseFunc(event);
                this.printToast(new ResponseData({}, stopStr, 201, true));
            };
        },
        stopWsConnect(wsName, closeStr, noConnectStr, killAll=false) {
            if (killAll) {
                this.killAll();
                return
            }
            const webSocket = this.webSockets[wsName];
            if (webSocket) {
                console.log(`Closing WebSocket connection for ${wsName}...`);
                webSocket.close();
                delete this.webSockets[wsName];
                this.printToast(new ResponseData({}, closeStr, 201,true))
            }
            else {
                console.warn(`No active WebSocket connection for ${wsName}.`);
                this.printToast(new ResponseData({}, noConnectStr, 0,false))
            }
        },
        killAll() {
            for (const wsName in this.webSockets) {
                const webSocket = this.webSockets[wsName];
                if (webSocket) {
                    console.log(`Closing WebSocket connection for ${wsName}...`);
                    // this.printToast(new ResponseData({}, closeStr, 201,true))
                    webSocket.close();
                    delete this.webSockets[wsName];
                }
            }
        },
    },
};
