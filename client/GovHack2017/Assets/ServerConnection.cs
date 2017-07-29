using System;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Text;
using UnityEngine;
using System.Linq;
using System.Collections.Generic;

// State object for receiving data from remote device.  
public class StateObject
{
    // Client socket.  
    public Socket workSocket = null;
    // Size of receive buffer.  
    public const int BufferSize = 256;
    // Receive buffer.  
    public byte[] buffer = new byte[BufferSize];
    // Received data string.  
    public StringBuilder sb = new StringBuilder();
}

public class ServerConnection : MonoBehaviour
{
    // The port number for the remote device.  
    private const int port = 23456;

    // The response from the remote device.  
    private static String response = String.Empty;

    private static Socket socket;

    private static void StartClient()
    {
        // Connect to a remote device.  
        try
        {
            // Establish the remote endpoint for the socket.  
            // The name of the   
            // remote device is "host.contoso.com".  
            

            IPHostEntry ipHostInfo = Dns.Resolve("localhost");
            IPAddress ipAddress = ipHostInfo
                .AddressList
                .Where(address => address.AddressFamily == AddressFamily.InterNetwork)
                .First();
            
            IPEndPoint remoteEP = new IPEndPoint(ipAddress, port);

            // Create a TCP/IP socket.  
            socket = new Socket(AddressFamily.InterNetwork,
                SocketType.Stream, ProtocolType.Tcp);

            // Connect to the remote endpoint.  
            socket.BeginConnect(remoteEP,
                new AsyncCallback(ConnectCallback), socket);
            
        }
        catch (Exception e)
        {
            Debug.LogFormat(e.ToString());
        }
    }

    private static void ConnectCallback(IAsyncResult ar)
    {
        try
        {
            // Retrieve the socket from the state object.  
            Socket client = (Socket)ar.AsyncState;

            // Complete the connection.  
            client.EndConnect(ar);

            Debug.LogFormat("Socket connected to {0}",
                client.RemoteEndPoint.ToString());

            Receive(client);
        }
        catch (Exception e)
        {
            Debug.LogFormat(e.ToString());
        }
    }

    private static void Receive(Socket client)
    {
        try
        {
            // Create the state object.  
            StateObject state = new StateObject();
            state.workSocket = client;

            // Begin receiving the data from the remote device.  
            client.BeginReceive(state.buffer, 0, StateObject.BufferSize, 0,
                new AsyncCallback(ReceiveCallback), state);
        }
        catch (Exception e)
        {
            Debug.LogFormat(e.ToString());
        }
    }

    public Queue<string> messages = new Queue<string>();

    private static void ReceiveCallback(IAsyncResult ar)
    {
        try
        {
            // Retrieve the state object and the client socket   
            // from the asynchronous state object.  
            StateObject state = (StateObject)ar.AsyncState;
            Socket client = state.workSocket;

            // Read data from the remote device.  
            int bytesRead = client.EndReceive(ar);

            if (bytesRead > 0)
            {

                var dataSoFar = Encoding.ASCII.GetString(state.buffer, 0, bytesRead);

                if (dataSoFar.Contains('\n'))
                {
                    // This is the end of the command.
                    // We might have more data after this, so split.

                    var splitData = new Queue<string>( dataSoFar.Split('\n'));

                    while (splitData.Count > 1)
                    {
                        var command = splitData.Dequeue();

                        state.sb.Append(command);
                        
                        MessageReceived(state.sb.ToString());

                        // Replace the stringbuilder
                        state.sb = new StringBuilder();
                    }
                    
                    // Put any data that wasn't part of a command on the buffer
                    state.sb.Append(splitData.Dequeue());
                    
                } else
                {
                    // There might be more data, so store the data received so far.  
                    state.sb.Append(dataSoFar);

                }


                // Get the rest of the data.  
                client.BeginReceive(state.buffer, 0, StateObject.BufferSize, 0,
                    new AsyncCallback(ReceiveCallback), state);
            }
        }
        catch (Exception e)
        {
            Debug.LogFormat(e.ToString());
        }
    }

    private static void Send(String data)
    {
        Debug.LogFormat("Sending: {0}", data);
        // Convert the string data to byte data using ASCII encoding.  
        byte[] byteData = Encoding.ASCII.GetBytes(data);

        // Begin sending the data to the remote device.  
        socket.BeginSend(byteData, 0, byteData.Length, 0,
            new AsyncCallback(SendCallback), socket);
    }

    private static void SendCallback(IAsyncResult ar)
    {
        try
        {
            // Retrieve the socket from the state object.  
            Socket client = (Socket)ar.AsyncState;

            // Complete sending the data to the remote device.  
            int bytesSent = client.EndSend(ar);
            //Debug.LogFormat("Sent {0} bytes to server.", bytesSent);
            
        }
        catch (Exception e)
        {
            Debug.LogFormat(e.ToString());
        }
    }

    public static event Action<string> MessageReceived;

    public void SendMessageToServer(string message)
    {
        Send(message + "\n");
    }

    private void Awake()
    {
        MessageReceived += ServerConnection_MessageReceived;
        StartClient();
    }

    private void ServerConnection_MessageReceived(string obj)
    {
        Debug.Log("Message received: " + obj);
    }

    public void SendNoOp()
    {
        SendMessageToServer("{}");
    }

    public void SendStartGame()
    {
        SendMessageToServer(@"{""action"": 1}");
    }
}