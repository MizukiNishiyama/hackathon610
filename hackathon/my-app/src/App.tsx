import './App.css';
import React, { useState, useEffect } from 'react';

//authentification
import { onAuthStateChanged } from "firebase/auth";
import { fireAuth } from "./firebase";
import { signInWithPopup, GoogleAuthProvider, signOut } from "firebase/auth";

// //routing
import { BrowserRouter, Route} from 'react-router-dom';

export const LoginForm: React.FC = () => {
    /**
     * googleでログインする
     */
    const signInWithGoogle = (): void => {
      // Google認証プロバイダを利用する
      const provider = new GoogleAuthProvider();
  
      // ログイン用のポップアップを表示
      signInWithPopup(fireAuth, provider)
        .then(res => {
            const user =res.user;
            alert("ログインユーザー: " + user.displayName);
            // fetch("http://localhost:3000/user", {
            fetch("https://uttc-bapgglyr6q-uc.a.run.app/user", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    name: user.displayName,
                    email: user.email
                })
            })    
        .then(response => response.json())
        .then(data => console.log(data))
        .catch((error) => {
            console.error('Error:', error);
        });
        })
        .catch(err => {
            const errorMessage =err.message;
            alert(errorMessage);
        })
    };
  
    /**
     * ログアウトする
     */
    const signOutWithGoogle = (): void => {
        signOut(fireAuth).then(() => {
          alert("ログアウトしました");
        }).catch(err => {
          alert(err);
        });
      };
    
  
    return (
      <div className = "loginform" >
        <button onClick={signInWithGoogle}>
          Googleでログイン
        </button>
        <button onClick={signOutWithGoogle}>
        ログアウト
        </button>
      </div>
    );
  };


type Channel = {
    id: string;
    name: string;
};

type Message = {
    messageid: string;
    content: string;
    userid: string;
    channelid: string;
    time: string;
};

type User = {
    name: string;
    email: string;
}

type Props = {
    activeChannel: string;
    setActiveChannel: React.Dispatch<React.SetStateAction<string>>;
    refreshMessages: boolean;
    setRefreshMessages: React.Dispatch<React.SetStateAction<boolean>>;
}

function ShowChannelMessage(props:Props) {
    const {activeChannel, setActiveChannel, refreshMessages} = props

    const [channels, setChannels] = useState<Channel[]>([]);
    const [messages, setMessages] = useState<Message[]>([]);

    useEffect(() => {
        const fetchChannels = async () => {
            // const response = await fetch("http://localhost:3000/getchannels");
            const response = await fetch('https://uttc-bapgglyr6q-uc.a.run.app/getchannels');
            const data = await response.json();
            console.log("表示するチャンネル",data);
            setChannels(data);
        };
        fetchChannels();
    }, []);

    useEffect(() => {
        const fetchMessages = async () => {
            if (activeChannel === "") {
                setMessages([]);
                return;
            }
            // const response = await fetch("http://localhost:3000/message?channelid=${activeChannel}");
            const response = await fetch(`https://uttc-bapgglyr6q-uc.a.run.app/message?channelid=${activeChannel}`);
            const data = await response.json();
            console.log("表示するメッセージ",data);
            setMessages(data);
            
        };

        fetchMessages();
    }, [activeChannel, refreshMessages]);
    

    return (
        <div className="showmessages">
            
            <div className="channels">
            <h1>channel</h1>
                {channels.map(channel => (
                    <div
                        key={channel.id}
                        onClick={() =>  setActiveChannel(channel.id)}
                        className={activeChannel === channel.id ? 'active' : ''}
                    >
                        {channel.name}
                    </div>
                ))}
            </div>
                
            <div className="messages">
            <h1>talk</h1>    
                {messages.map(message => (
                    <div key={message.messageid}>
                        <span className="user-name">{message.userid}</span>
                        <div className ="content-time">
                            <span className="message-content">{message.content}</span>
                            <span className="message-time">{message.time}</span>
                        </div>
                    </div>
                    
                ))}
            </div>
        </div>
    );
}

// function getUserinfo () {
    

// }

function Sendmessage(props:Props) {
    const {activeChannel, refreshMessages, setRefreshMessages} = props
    const [content, setContent] = useState("")
    const [userid ,setUserid] = useState("")
    
    
    
    const sendMessages = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log(content);
        setContent("")

        if (!content) {
            alert("メッセージを入力してください。");
            return;
        }

        try {
            const user = fireAuth.currentUser
            if (!user) {
                alert("ユーザーがログインしていません。");
                return;
            }
            //const response = await fetch("http://localhost:3000/message")
            const response = await fetch("https://uttc-bapgglyr6q-uc.a.run.app/message", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    content: content,
                    channelid: activeChannel,
                    userid: user.displayName,
                    time: ` (${new Date().getMonth()+1}/${new Date().getDate()} ${new Date().getHours()}:${new Date().getMinutes()})`,
                }),
                
            });

            const data = await response.json();
            console.log("success", data);
            setRefreshMessages(!refreshMessages)
        } catch (error) {
            console.error("error:", error);
        }
    
    };
 
    return(
        <div className="sendmessages">
                <form onSubmit={sendMessages}>
                    <h1>
                        MESSAGE 
                        <input type="text" value={content} onChange={(e) => setContent(e.target.value)} />
                    </h1>  
                    <button type ="submit">SEND</button>
                </form>
            </div>
    )
}

// function LoginState() {
//     const [loginUser, setLoginUser] = useState(fireAuth.currentUser);
  
//   // ログイン状態を監視して、stateをリアルタイムで更新する
//     onAuthStateChanged(fireAuth, user => {
//         setLoginUser(user);
//     });
//     return(
//         <div>
//             <LoginForm />
//             {loginUser ? <App /> : null}
//         </div>
        
//     );
// }


function App() {
    const [activeChannel, setActiveChannel] = useState<string>("");
    const [refreshMessages, setRefreshMessages] = useState<boolean>(false);
    return (
        // <BrowserRouter>
                <div className="App">
                    <LoginForm /> 
                    <ShowChannelMessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} setRefreshMessages={setRefreshMessages} refreshMessages={refreshMessages}/>
                    <Sendmessage activeChannel={activeChannel} setActiveChannel={setActiveChannel} setRefreshMessages={setRefreshMessages} refreshMessages={refreshMessages}/>     
                                
                </div>
        // </BrowserRouter>
    );
}

export default App;


