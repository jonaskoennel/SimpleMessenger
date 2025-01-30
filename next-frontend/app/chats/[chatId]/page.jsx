import ChatBar from './chatbox';

export default async function ChatPage({params}) {
  console.log(params.chatId)
    return (
      <div className="flex flex-row h-screen">
        <ChatBar />
        <div className="basis-3/4 bg-slate-600">
            Test
        </div>
      </div>
    );
}