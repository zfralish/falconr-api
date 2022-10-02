CREATE TABLE if not exists birds (
  id uuid primary key,
  falconer_id text not null,
  name text not null,
  color text not null,
  species text not null,
  trap_date timestamp with time zone not null default now(),
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);

create table if not exists weights (
  id uuid primary key,
  bird_id uuid references birds not null,
  weight float not null,
  w_time timestamp with time zone not null default now(),
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);

create table if not exists feedings (
  id uuid primary key,
  bird_id uuid references birds not null,
  f_time timestamp with time zone not null default now(),
  food_type text not null,
  amount float not null,
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);

create table if not exists hunts (
  id uuid primary key,
  bird_id uuid references birds not null,
  start_time timestamp with time zone not null default now(),
  end_time timestamp with time zone not null default now(),
  prey_type text not null,
  prey_count int not null,
  notes text,
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);

create table if not exists trainings (
  id uuid primary key,
  bird_id uuid references birds not null,
  start_time timestamp with time zone not null default now(),
  end_time timestamp with time zone not null default now(),
  training_type text not null,
  notes text,
  performance int not null default 5,
  created_at timestamp with time zone default now(),
  updated_at timestamp with time zone default now()
);