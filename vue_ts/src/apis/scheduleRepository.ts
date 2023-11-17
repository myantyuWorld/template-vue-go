import { AxiosPromise } from "axios";
import repository from "./repository";
import { Schedule } from "./petRepository";

const resource = "/pet";

export default class ScheduleRepository {
  public create(id: number, payload: Schedule): AxiosPromise<Schedule> {
    return repository.post(`${resource}/${id}}/schedule`, payload);
  }
}
