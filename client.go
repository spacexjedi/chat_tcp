package main

import (
"bufio"
"fmt"
"net"
"string"
)

type client struct {
  conn net.Conn
  nick string
  room *room
  commands chan<- command
}

func (c *client) readInput() {
  
  for {

    msg, err := bufio.NewReader(c.conn).ReadString('\n')
    if err != nill { // diferente
      return
    }

    msg = strings.Trim(msg, "\r\n")
    args := strings.Split(msg, " ")
    cmd =: strings.TrimSpace(args[0])

    switch cmd {
    case "/nick":
      c.commands <- command {

        id: CMD_NICK,
        client: c,
        args: args,

      } // end of case 1

      case "/join":
        c.commands <- command {

          id: CMD_JOIN,
          client: c,
          args: args,

        } // end of case 2

        case "/rooms":
          c.commands <- command {
            id: CMD_ROOMS,
            client: c,
          } // end of case 3

          case "/msg":
            c.commands <- command {
              id: CMD_MSG,
              client: c,
              args: args,
            } // end of case 4

            case "/leave":
              c.commands <- command {
                id: CMD_LEAVE,
                client: c,
              } // end of case 5

              default:
                c.err(fmt.Errorf("Unknow command %s", cmd))
      
    } // end of switch




  } // end of for

}

func (c *client) err(err error) {
  c.conn.Write([]byte("err: " + err.Error() + "\n"))
}


func (c *client) msg(msg string) {

	c.conn.Write([]byte("> " + msg + "\n"))
	
}