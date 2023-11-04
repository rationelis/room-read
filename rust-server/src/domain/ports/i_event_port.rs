use crate::domain::model::event::Event;

pub trait EventPort {
    fn GetEvents(&self) -> Vec<Event>;
}
