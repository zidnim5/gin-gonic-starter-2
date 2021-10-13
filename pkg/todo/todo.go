package todo

import (
     conn "starter/internal/postgresdb"
     "time"
)


type Todo struct{
     Name      string    `json:"name"`
     CreatedAt time.Time `json:"created_at"`
}

func (td *Todo) TableName() string{
     return "todo"
}

func (td *Todo) FindId(id string) (err error){
     con, _ := conn.DB.DB()
     defer con.Close()
     
     if err = conn.DB.Where("id =?", id).Find(td).Error; err != nil {
          return err
     }
     
     return nil
}

func (td *Todo) Insert() (err error){
     con, _ := conn.DB.DB()
     defer con.Close()
     
     if err = conn.DB.Create(td).Error; err != nil {
          return err
     }
     return nil
}