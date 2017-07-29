using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Net.Sockets;
using System.Threading;


public class ServerConnection : MonoBehaviour {


    class Connection
    {

        private static Thread thread;

        private static Queue<string> outgoingMessages = new Queue<string>();
        private static Queue<string> incomingMessages = new Queue<string>();

        public static void PutMessage(string message)
        {
            Monitor.Enter(outgoingMessages);

            outgoingMessages.Enqueue(message);

            Monitor.Exit(outgoingMessages);
        }

        public static string GetMessage()
        {
            if (Monitor.TryEnter(incomingMessages))
            {
                try
                {
                    if (incomingMessages.Count > 0)
                    {
                        return incomingMessages.Dequeue();
                    }
                    else
                    {
                        return null;
                    }
                } finally {
                    Monitor.Exit(incomingMessages);
                }
            }
            return null;
            
        }

        private static void RunCommunication()
        {

            PutMessage("Connecting to " + host + ":" + port);
            var client = new TcpClient(host, port);

            var stream = client.GetStream();

            PutMessage("Connected to " + host + ":" + port);

            var incomingBytes = new List<byte>();

            while (true)
            {
                // Deliver the next outgoing message
                Monitor.Enter(outgoingMessages);

                if (outgoingMessages.Count > 0)
                {
                    var nextOutgoingMessage = outgoingMessages.Dequeue();

                    var outBytes = System.Text.Encoding.UTF8.GetBytes(nextOutgoingMessage);

                    stream.Write(outBytes, 0, outBytes.Length);
                }

                Monitor.Exit(outgoingMessages);

                // Read data; add data to the queue if able
                if (stream.DataAvailable)
                {
                    var nextByte = (byte)stream.ReadByte();

                    if (nextByte == '\n')
                    {
                        var bytes = incomingBytes.ToArray();

                        var message = System.Text.Encoding.UTF8.GetString(bytes);

                        PutMessage(message);

                        incomingBytes.Clear();
                    } else
                    {
                        incomingBytes.Add(nextByte);
                    }
                }              
            }
            
        }

        private static string host;
        private static int port;

        public static void Connect(string host, int port)
        {
            if (thread != null)
            {
                thread.Abort();
            }
            Connection.host = host;
            Connection.port = port;

            thread = new Thread(new ThreadStart(RunCommunication));
            thread.IsBackground = true;
            thread.Start();
        }
    }
    
    public string host = "localhost";
    public int port = 23456;

    public void Awake()
    {
        Connection.Connect(host, port);
    }

    // Update is called once per frame
    void Update () {
		
        var message = Connection.GetMessage();

        if (message != null)
        {
            Debug.LogFormat("Message: {0}", message);
        }
	}

    public void SendMessageToServer(string message)
    {
        Connection.PutMessage(message);
    }
}
