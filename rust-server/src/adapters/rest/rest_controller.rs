use std::convert::Infallible;

use async_trait::async_trait;
use hyper::{Body, Request, Response};

use crate::{
    domain::ports::i_event_port::EventPort,
    infrastructure::configuration::configuration::Configuration,
};

#[async_trait]
trait RestController {
    async fn GetEvents(
        &self,
        req: Request<Body>,
    ) -> std::result::Result<Response<Body>, Infallible>;
}

struct RestControllerImpl {
    configuration: Configuration,
    event_port: dyn EventPort,
}

pub fn new_rest_controller(
    configuration: Configuration,
    event_port: dyn EventPort,
) -> RestControllerImpl {
    RestControllerImpl {
        configuration,
        event_port,
    }
}

impl RestController for RestControllerImpl {
    async fn responses(req: Request<Body>) -> std::result::Result<Response<Body>, Infallible> {
        match (req.method(), req.uri().path()) {
            (&Method::GET, "/event") => get_events(req),
            _ => Ok(Response::builder()
                .status(StatusCode::NOT_FOUND)
                .body(NOTFOUND.into())
                .unwrap()),
        }
    }

    fn get_events(&self, req: Request<Body>) -> std::result::Result<Response<Body>, Infallible> {
        let events = self.event_port.GetEvents();
        Ok(Response::new(events.into()))
    }
}
