package handler

// /events?filter=day&date=2021-10-10
// /events?filter=week&date=2021-10-10
// /events?filter=month&date=2021-10-10

//type GetEventsRequest struct {
//}
//
//type GetEventsResponse struct {
//}
//
//func GetEventsHandler() func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		filter := r.URL.Query().Get("filter")
//		date := r.URL.Query().Get("date")
//
//		var events []storage.Event
//		if filter == "day" {
//			events = getEventsByDay(app, date)
//		} else if filter == "week" {
//			events = getEventsByWeek(app, date)
//		} else if filter == "month" {
//			events = getEventsByMonth(app, date)
//		} else {
//			http.Error(w, "invalid filter", http.StatusBadRequest)
//			return
//		}
//
//		// write response
//		w.WriteHeader(http.StatusOK)
//	}
//}
//
//func getEventsByDay(app internalhttp.Application, date string) []storage.Event {
//	return []storage.Event{}
//}
//
//func getEventsByWeek(app internalhttp.Application, date string) []storage.Event {
//	return []storage.Event{}
//}
//
//func getEventsByMonth(app internalhttp.Application, date string) []storage.Event {
//	return []storage.Event{}
//}
