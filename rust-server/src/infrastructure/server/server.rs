use hyper::client::HttpConnector;
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Client, Method, Request, Response, Result, Server, StatusCode};

use crate::infrastructure::configuration::configuration::Configuration;

type GenericError = Box<dyn std::error::Error + Send + Sync>;

static NOTFOUND: &[u8] = b"Not Found";
static INDEX: &[u8] = b"<a href=\"test.html\">test.html</a>";

pub struct RoomReadServer {
    pub configuration: Configuration,
    // pub server: Server<AddrIncoming>,
}

async fn response_examples(
    req: Request<Body>,
    client: Client<HttpConnector>,
) -> Result<Response<Body>> {
    match (req.method(), req.uri().path()) {
        (&Method::GET, "/") | (&Method::GET, "/index.html") => Ok(Response::new(INDEX.into())),
        // (&Method::GET, "/event") => get
        _ => Ok(Response::builder()
            .status(StatusCode::NOT_FOUND)
            .body(NOTFOUND.into())
            .unwrap()),
    }
}

impl RoomReadServer {
    pub fn new(configuration: Configuration) -> RoomReadServer {
        let addr = format!(
            "{}:{}",
            configuration.server.host, configuration.server.port
        )
        .parse()
        .unwrap();

        // let client = Client::new();

        let restController = new_rest_controller(configuration.clone(), event_port);

        let new_service = make_service_fn(move |_| {
            // let client = client.clone();
            async {
                Ok::<_, GenericError>(service_fn(move |req| {
                    response_examples(req, client.to_owned())
                }))
            }
        });

        let server = Server::bind(&addr).serve(new_service);

        RoomReadServer {
            configuration,
            // server,
        }
    }
}
